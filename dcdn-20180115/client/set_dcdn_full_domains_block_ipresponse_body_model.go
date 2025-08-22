// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnFullDomainsBlockIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SetDcdnFullDomainsBlockIPResponseBody
	GetCode() *int32
	SetMessage(v string) *SetDcdnFullDomainsBlockIPResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetDcdnFullDomainsBlockIPResponseBody
	GetRequestId() *string
}

type SetDcdnFullDomainsBlockIPResponseBody struct {
	// The response code.
	//
	// If the value of Code is not 0, specific required parameters are missing or the parameter format is invalid.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-802B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDcdnFullDomainsBlockIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnFullDomainsBlockIPResponseBody) GoString() string {
	return s.String()
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) SetCode(v int32) *SetDcdnFullDomainsBlockIPResponseBody {
	s.Code = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) SetMessage(v string) *SetDcdnFullDomainsBlockIPResponseBody {
	s.Message = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) SetRequestId(v string) *SetDcdnFullDomainsBlockIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDcdnFullDomainsBlockIPResponseBody) Validate() error {
	return dara.Validate(s)
}
