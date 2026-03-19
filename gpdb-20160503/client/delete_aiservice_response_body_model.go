// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAIServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAIServiceResponseBody
	GetRequestId() *string
}

type DeleteAIServiceResponseBody struct {
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAIServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAIServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAIServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAIServiceResponseBody) SetRequestId(v string) *DeleteAIServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAIServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
