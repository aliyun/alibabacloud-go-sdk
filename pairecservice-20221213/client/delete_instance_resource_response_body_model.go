// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInstanceResourceResponseBody
	GetRequestId() *string
}

type DeleteInstanceResourceResponseBody struct {
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResourceResponseBody) SetRequestId(v string) *DeleteInstanceResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
