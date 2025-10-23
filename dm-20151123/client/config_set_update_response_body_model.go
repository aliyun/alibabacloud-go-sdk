// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetUpdateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ConfigSetUpdateResponseBody
	GetId() *string
	SetRequestId(v string) *ConfigSetUpdateResponseBody
	GetRequestId() *string
}

type ConfigSetUpdateResponseBody struct {
	// example:
	//
	// xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSetUpdateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetUpdateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetUpdateResponseBody) GetId() *string {
	return s.Id
}

func (s *ConfigSetUpdateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetUpdateResponseBody) SetId(v string) *ConfigSetUpdateResponseBody {
	s.Id = &v
	return s
}

func (s *ConfigSetUpdateResponseBody) SetRequestId(v string) *ConfigSetUpdateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetUpdateResponseBody) Validate() error {
	return dara.Validate(s)
}
