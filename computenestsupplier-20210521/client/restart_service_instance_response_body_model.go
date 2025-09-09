// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartServiceInstanceResponseBody
	GetRequestId() *string
}

type RestartServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartServiceInstanceResponseBody) SetRequestId(v string) *RestartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
