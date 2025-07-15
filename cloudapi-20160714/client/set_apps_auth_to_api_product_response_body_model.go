// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppsAuthToApiProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAppsAuthToApiProductResponseBody
	GetRequestId() *string
}

type SetAppsAuthToApiProductResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAppsAuthToApiProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAppsAuthToApiProductResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppsAuthToApiProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAppsAuthToApiProductResponseBody) SetRequestId(v string) *SetAppsAuthToApiProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAppsAuthToApiProductResponseBody) Validate() error {
	return dara.Validate(s)
}
