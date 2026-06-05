// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComfyProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteComfyProductionResponseBody
	GetCode() *int64
	SetMessage(v string) *DeleteComfyProductionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteComfyProductionResponseBody
	GetRequestId() *string
}

type DeleteComfyProductionResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteComfyProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComfyProductionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComfyProductionResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteComfyProductionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteComfyProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComfyProductionResponseBody) SetCode(v int64) *DeleteComfyProductionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteComfyProductionResponseBody) SetMessage(v string) *DeleteComfyProductionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteComfyProductionResponseBody) SetRequestId(v string) *DeleteComfyProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComfyProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
