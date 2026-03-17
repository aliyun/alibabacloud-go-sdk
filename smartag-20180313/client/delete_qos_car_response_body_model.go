// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosCarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQosCarResponseBody
	GetRequestId() *string
}

type DeleteQosCarResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 551CD836-9E46-4F2C-A167-B4363180A647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQosCarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosCarResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQosCarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQosCarResponseBody) SetRequestId(v string) *DeleteQosCarResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQosCarResponseBody) Validate() error {
	return dara.Validate(s)
}
