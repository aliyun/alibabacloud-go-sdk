// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSampleApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadSampleApiResponseBody
	GetCode() *string
	SetMessage(v string) *UploadSampleApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadSampleApiResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UploadSampleApiResponseBody
	GetSuccess() *string
}

type UploadSampleApiResponseBody struct {
	// Request code returned
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message returned
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicator of whether the access was successful
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadSampleApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadSampleApiResponseBody) GoString() string {
	return s.String()
}

func (s *UploadSampleApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadSampleApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadSampleApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadSampleApiResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UploadSampleApiResponseBody) SetCode(v string) *UploadSampleApiResponseBody {
	s.Code = &v
	return s
}

func (s *UploadSampleApiResponseBody) SetMessage(v string) *UploadSampleApiResponseBody {
	s.Message = &v
	return s
}

func (s *UploadSampleApiResponseBody) SetRequestId(v string) *UploadSampleApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadSampleApiResponseBody) SetSuccess(v string) *UploadSampleApiResponseBody {
	s.Success = &v
	return s
}

func (s *UploadSampleApiResponseBody) Validate() error {
	return dara.Validate(s)
}
