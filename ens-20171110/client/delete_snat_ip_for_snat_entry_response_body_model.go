// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatIpForSnatEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSnatIpForSnatEntryResponseBody
	GetRequestId() *string
}

type DeleteSnatIpForSnatEntryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnatIpForSnatEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatIpForSnatEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnatIpForSnatEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSnatIpForSnatEntryResponseBody) SetRequestId(v string) *DeleteSnatIpForSnatEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSnatIpForSnatEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
