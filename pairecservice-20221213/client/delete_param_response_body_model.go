// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParamResponseBody
	GetRequestId() *string
}

type DeleteParamResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F0AB6527-093F-5C44-B3BD-42C8C210C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParamResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParamResponseBody) SetRequestId(v string) *DeleteParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParamResponseBody) Validate() error {
	return dara.Validate(s)
}
