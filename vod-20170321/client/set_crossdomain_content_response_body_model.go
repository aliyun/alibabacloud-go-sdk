// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCrossdomainContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetCrossdomainContentResponseBody
	GetRequestId() *string
}

type SetCrossdomainContentResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A13-****-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCrossdomainContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCrossdomainContentResponseBody) GoString() string {
	return s.String()
}

func (s *SetCrossdomainContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCrossdomainContentResponseBody) SetRequestId(v string) *SetCrossdomainContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCrossdomainContentResponseBody) Validate() error {
	return dara.Validate(s)
}
