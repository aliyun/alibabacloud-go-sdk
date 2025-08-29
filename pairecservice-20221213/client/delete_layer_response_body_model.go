// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLayerResponseBody
	GetRequestId() *string
}

type DeleteLayerResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F457D79-C4A2-5E8C-83E4-0D089456E2AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLayerResponseBody) SetRequestId(v string) *DeleteLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLayerResponseBody) Validate() error {
	return dara.Validate(s)
}
