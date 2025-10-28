// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ContinuePipelineResponseBody
	GetCode() *int32
	SetMessage(v string) *ContinuePipelineResponseBody
	GetMessage() *string
	SetRequestId(v string) *ContinuePipelineResponseBody
	GetRequestId() *string
}

type ContinuePipelineResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D16979DC-4D42-****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinuePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuePipelineResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ContinuePipelineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ContinuePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinuePipelineResponseBody) SetCode(v int32) *ContinuePipelineResponseBody {
	s.Code = &v
	return s
}

func (s *ContinuePipelineResponseBody) SetMessage(v string) *ContinuePipelineResponseBody {
	s.Message = &v
	return s
}

func (s *ContinuePipelineResponseBody) SetRequestId(v string) *ContinuePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinuePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
