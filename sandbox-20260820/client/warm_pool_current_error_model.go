// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWarmPoolCurrentError interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *WarmPoolCurrentError
	GetCode() *string
	SetMessage(v string) *WarmPoolCurrentError
	GetMessage() *string
}

type WarmPoolCurrentError struct {
	// The error code of the current error.
	//
	// example:
	//
	// SourceNotFound
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the current error.
	//
	// example:
	//
	// Source template does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s WarmPoolCurrentError) String() string {
	return dara.Prettify(s)
}

func (s WarmPoolCurrentError) GoString() string {
	return s.String()
}

func (s *WarmPoolCurrentError) GetCode() *string {
	return s.Code
}

func (s *WarmPoolCurrentError) GetMessage() *string {
	return s.Message
}

func (s *WarmPoolCurrentError) SetCode(v string) *WarmPoolCurrentError {
	s.Code = &v
	return s
}

func (s *WarmPoolCurrentError) SetMessage(v string) *WarmPoolCurrentError {
	s.Message = &v
	return s
}

func (s *WarmPoolCurrentError) Validate() error {
	return dara.Validate(s)
}
