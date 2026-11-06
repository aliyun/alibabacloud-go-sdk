// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoDisposeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAutoDisposeConfigResponseBody
	GetRequestId() *string
}

type CreateAutoDisposeConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6D7B26AD-2EDD-5A96-B9D2-CDEABFC55690
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAutoDisposeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoDisposeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoDisposeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoDisposeConfigResponseBody) SetRequestId(v string) *CreateAutoDisposeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoDisposeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
