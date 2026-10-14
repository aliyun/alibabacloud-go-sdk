// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNetworkChannelResponseBody
	GetRequestId() *string
}

type CreateNetworkChannelResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F9F1CB1A-B1D5-4EF5-A53A-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkChannelResponseBody) SetRequestId(v string) *CreateNetworkChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
