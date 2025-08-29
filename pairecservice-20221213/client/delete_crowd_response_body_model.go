// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrowdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCrowdResponseBody
	GetRequestId() *string
}

type DeleteCrowdResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A04CB8C0-E74A-5E83-BC61-64D153574EC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCrowdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrowdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCrowdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCrowdResponseBody) SetRequestId(v string) *DeleteCrowdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCrowdResponseBody) Validate() error {
	return dara.Validate(s)
}
