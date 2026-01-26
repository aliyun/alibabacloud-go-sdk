// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomHostnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostname(v string) *CreateCustomHostnameResponseBody
	GetHostname() *string
	SetHostnameId(v int64) *CreateCustomHostnameResponseBody
	GetHostnameId() *int64
	SetRequestId(v string) *CreateCustomHostnameResponseBody
	GetRequestId() *string
}

type CreateCustomHostnameResponseBody struct {
	// example:
	//
	// custom.site.com
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 1234567890123
	HostnameId *int64 `json:"HostnameId,omitempty" xml:"HostnameId,omitempty"`
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomHostnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomHostnameResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomHostnameResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *CreateCustomHostnameResponseBody) GetHostnameId() *int64 {
	return s.HostnameId
}

func (s *CreateCustomHostnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomHostnameResponseBody) SetHostname(v string) *CreateCustomHostnameResponseBody {
	s.Hostname = &v
	return s
}

func (s *CreateCustomHostnameResponseBody) SetHostnameId(v int64) *CreateCustomHostnameResponseBody {
	s.HostnameId = &v
	return s
}

func (s *CreateCustomHostnameResponseBody) SetRequestId(v string) *CreateCustomHostnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomHostnameResponseBody) Validate() error {
	return dara.Validate(s)
}
