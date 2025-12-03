// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGlobalResourceResponseBody
	GetRequestId() *string
}

type DeleteGlobalResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGlobalResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGlobalResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGlobalResourceResponseBody) SetRequestId(v string) *DeleteGlobalResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGlobalResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
