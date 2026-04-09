// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteDomainResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DeleteDomainResponseBody
	GetRequestId() *string
}

type DeleteDomainResponseBody struct {
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// FA8C2599-362D-4113-8FB4-E88A40C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDomainResponseBody) SetDomainName(v string) *DeleteDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
