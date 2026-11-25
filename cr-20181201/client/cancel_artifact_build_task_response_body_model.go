// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelArtifactBuildTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelArtifactBuildTaskResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *CancelArtifactBuildTaskResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *CancelArtifactBuildTaskResponseBody
	GetRequestId() *string
}

type CancelArtifactBuildTaskResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API call was successful. Valid values:
	//
	// - `true`: The API call was successful.
	//
	// - `false`: The API call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelArtifactBuildTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelArtifactBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelArtifactBuildTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelArtifactBuildTaskResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *CancelArtifactBuildTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelArtifactBuildTaskResponseBody) SetCode(v string) *CancelArtifactBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelArtifactBuildTaskResponseBody) SetIsSuccess(v bool) *CancelArtifactBuildTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CancelArtifactBuildTaskResponseBody) SetRequestId(v string) *CancelArtifactBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelArtifactBuildTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
