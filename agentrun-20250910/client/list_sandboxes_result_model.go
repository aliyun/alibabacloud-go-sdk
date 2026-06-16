// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSandboxesResult
	GetCode() *string
	SetData(v *ListSandboxesOutput) *ListSandboxesResult
	GetData() *ListSandboxesOutput
	SetRequestId(v string) *ListSandboxesResult
	GetRequestId() *string
}

type ListSandboxesResult struct {
	// The status of the request. A value of \\"SUCCESS\\" indicates that the request was successful. If the request fails, an error code is returned, such as \\"ERR_BAD_REQUEST\\", \\"ERR_VALIDATION_FAILED\\", or \\"ERR_INTERNAL_SERVER_ERROR\\".
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Contains the list of sandboxes.
	Data *ListSandboxesOutput `json:"data,omitempty" xml:"data,omitempty"`
	// The unique request ID used for issue tracking.
	//
	// example:
	//
	// 55D4BE40-2811-5CFB-8482-E0E98D575B1E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSandboxesResult) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesResult) GoString() string {
	return s.String()
}

func (s *ListSandboxesResult) GetCode() *string {
	return s.Code
}

func (s *ListSandboxesResult) GetData() *ListSandboxesOutput {
	return s.Data
}

func (s *ListSandboxesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSandboxesResult) SetCode(v string) *ListSandboxesResult {
	s.Code = &v
	return s
}

func (s *ListSandboxesResult) SetData(v *ListSandboxesOutput) *ListSandboxesResult {
	s.Data = v
	return s
}

func (s *ListSandboxesResult) SetRequestId(v string) *ListSandboxesResult {
	s.RequestId = &v
	return s
}

func (s *ListSandboxesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
