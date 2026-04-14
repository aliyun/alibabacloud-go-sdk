// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ConfigSetCreateResponseBody
	GetId() *string
	SetRequestId(v string) *ConfigSetCreateResponseBody
	GetRequestId() *string
}

type ConfigSetCreateResponseBody struct {
	// example:
	//
	// XXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// XXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSetCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCreateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetCreateResponseBody) GetId() *string {
	return s.Id
}

func (s *ConfigSetCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetCreateResponseBody) SetId(v string) *ConfigSetCreateResponseBody {
	s.Id = &v
	return s
}

func (s *ConfigSetCreateResponseBody) SetRequestId(v string) *ConfigSetCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
