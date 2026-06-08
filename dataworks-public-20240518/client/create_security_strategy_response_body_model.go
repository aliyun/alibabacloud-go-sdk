// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateSecurityStrategyResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateSecurityStrategyResponseBody
	GetRequestId() *string
}

type CreateSecurityStrategyResponseBody struct {
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSecurityStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityStrategyResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateSecurityStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityStrategyResponseBody) SetId(v int64) *CreateSecurityStrategyResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSecurityStrategyResponseBody) SetRequestId(v string) *CreateSecurityStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
