// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePipelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePipelineResponseBody
	GetRequestId() *string
}

type CreatePipelineResponseBody struct {
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreatePipelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePipelineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePipelineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePipelineResponseBody) SetRequestId(v string) *CreatePipelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePipelineResponseBody) Validate() error {
	return dara.Validate(s)
}
