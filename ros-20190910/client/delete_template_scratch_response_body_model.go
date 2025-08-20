// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateScratchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTemplateScratchResponseBody
	GetRequestId() *string
}

type DeleteTemplateScratchResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1B9C6333-5904-5A1E-9845-CB17A369AFDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateScratchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateScratchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateScratchResponseBody) SetRequestId(v string) *DeleteTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateScratchResponseBody) Validate() error {
	return dara.Validate(s)
}
