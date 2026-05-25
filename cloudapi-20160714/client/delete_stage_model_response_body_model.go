// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStageModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStageModelResponseBody
	GetRequestId() *string
}

type DeleteStageModelResponseBody struct {
	// example:
	//
	// CEB6EC62-B6C7-5082-A45A-45A204724xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStageModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStageModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStageModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStageModelResponseBody) SetRequestId(v string) *DeleteStageModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStageModelResponseBody) Validate() error {
	return dara.Validate(s)
}
