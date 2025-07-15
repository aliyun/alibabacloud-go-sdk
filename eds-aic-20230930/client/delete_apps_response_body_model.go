// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppsResponseBody
	GetRequestId() *string
}

type DeleteAppsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 310A783E-CC46-5452-A8A3-71AE5DB5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppsResponseBody) SetRequestId(v string) *DeleteAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppsResponseBody) Validate() error {
	return dara.Validate(s)
}
