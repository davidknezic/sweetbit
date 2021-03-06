import React from 'react';
import ReactModal from 'react-modal';
import css from 'styled-jsx/css';

ReactModal.setAppElement('#root')

const { className, styles } = css.resolve`
  .ReactModal__Overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    overflow-y: auto;
    opacity: 0;
    transition: opacity 200ms ease-in-out;
    background: rgba(0, 0, 0, 0.6);
    scroll-behavior: smooth;
  }
  .ReactModal__Overlay--after-open {
    opacity: 1;
  }
  .ReactModal__Overlay--before-close {
    opacity: 0;
  }
  .ReactModal__Content {
    max-width: 460px;
    width: 100%;
    margin: 0 auto;
    border-radius: none;
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    min-height: 100vh;
    background: white;
  }
  .ReactModal__Content > :global(div) {
    flex: 1;
  }
  @media (min-width: 500px) {
    .ReactModal__Content {
      display: block;
      border-radius: 10px;
      overflow: hidden;
      max-width: 500px;
      margin: 20px auto;
      min-height: auto;
    }
  }
`

function Modal({ children, open, onClose }) {
  return (
    <ReactModal
      closeTimeoutMS={200}
      isOpen={open}
      className={className}
      overlayClassName={className}
      contentLabel="modal"
      onRequestClose={onClose}
    >
      {children}
      {styles}
    </ReactModal>
  );
}

export default Modal;
