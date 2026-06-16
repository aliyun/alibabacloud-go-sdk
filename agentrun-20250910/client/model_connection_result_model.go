// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConnectionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelConnectionResult
	GetCode() *string
	SetData(v *ModelConnection) *ModelConnectionResult
	GetData() *ModelConnection
	SetRequestId(v string) *ModelConnectionResult
	GetRequestId() *string
}

type ModelConnectionResult struct {
	// `SUCCESS` indicates a successful request. On failure, this field returns the error type.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The resulting model connection object.
	//
	// example:
	//
	// {}
	Data *ModelConnection `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request ID used for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ModelConnectionResult) String() string {
	return dara.Prettify(s)
}

func (s ModelConnectionResult) GoString() string {
	return s.String()
}

func (s *ModelConnectionResult) GetCode() *string {
	return s.Code
}

func (s *ModelConnectionResult) GetData() *ModelConnection {
	return s.Data
}

func (s *ModelConnectionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelConnectionResult) SetCode(v string) *ModelConnectionResult {
	s.Code = &v
	return s
}

func (s *ModelConnectionResult) SetData(v *ModelConnection) *ModelConnectionResult {
	s.Data = v
	return s
}

func (s *ModelConnectionResult) SetRequestId(v string) *ModelConnectionResult {
	s.RequestId = &v
	return s
}

func (s *ModelConnectionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
