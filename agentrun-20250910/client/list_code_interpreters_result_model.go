// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpretersResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCodeInterpretersResult
	GetCode() *string
	SetData(v *ListCodeInterpretersOutput) *ListCodeInterpretersResult
	GetData() *ListCodeInterpretersOutput
	SetRequestId(v string) *ListCodeInterpretersResult
	GetRequestId() *string
}

type ListCodeInterpretersResult struct {
	// The response status code. `SUCCESS` indicates the request was successful. If the request fails, this field contains an error type, such as `ERR_BAD_REQUEST`, `ERR_VALIDATION_FAILED`, or `ERR_INTERNAL_SERVER_ERROR`.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// A list of code interpreter objects.
	//
	// example:
	//
	// {}
	Data *ListCodeInterpretersOutput `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request ID, used for troubleshooting.
	//
	// example:
	//
	// F8A0F5F3-0C3E-4C82-9D4F-5E4B6A7C8D9E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCodeInterpretersResult) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpretersResult) GoString() string {
	return s.String()
}

func (s *ListCodeInterpretersResult) GetCode() *string {
	return s.Code
}

func (s *ListCodeInterpretersResult) GetData() *ListCodeInterpretersOutput {
	return s.Data
}

func (s *ListCodeInterpretersResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCodeInterpretersResult) SetCode(v string) *ListCodeInterpretersResult {
	s.Code = &v
	return s
}

func (s *ListCodeInterpretersResult) SetData(v *ListCodeInterpretersOutput) *ListCodeInterpretersResult {
	s.Data = v
	return s
}

func (s *ListCodeInterpretersResult) SetRequestId(v string) *ListCodeInterpretersResult {
	s.RequestId = &v
	return s
}

func (s *ListCodeInterpretersResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
