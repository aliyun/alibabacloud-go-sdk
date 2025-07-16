// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCdnFullDomainsBlockIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SetCdnFullDomainsBlockIPResponseBody
	GetCode() *int32
	SetMessage(v string) *SetCdnFullDomainsBlockIPResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetCdnFullDomainsBlockIPResponseBody
	GetRequestId() *string
}

type SetCdnFullDomainsBlockIPResponseBody struct {
	// The status code. The status code 0 indicates that the call is successful. If another status code is returned, the call fails.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned. If the request was successful, OK is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 23ACE7E2-2302-42E3-98F8-E5E697FD86C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetCdnFullDomainsBlockIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetCdnFullDomainsBlockIPResponseBody) GoString() string {
	return s.String()
}

func (s *SetCdnFullDomainsBlockIPResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SetCdnFullDomainsBlockIPResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetCdnFullDomainsBlockIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetCdnFullDomainsBlockIPResponseBody) SetCode(v int32) *SetCdnFullDomainsBlockIPResponseBody {
	s.Code = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPResponseBody) SetMessage(v string) *SetCdnFullDomainsBlockIPResponseBody {
	s.Message = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPResponseBody) SetRequestId(v string) *SetCdnFullDomainsBlockIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetCdnFullDomainsBlockIPResponseBody) Validate() error {
	return dara.Validate(s)
}
