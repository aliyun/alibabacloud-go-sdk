// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileProtectClientRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateFileProtectClientRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateFileProtectClientRuleResponseBody
	GetRequestId() *string
}

type CreateFileProtectClientRuleResponseBody struct {
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BREF20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileProtectClientRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileProtectClientRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileProtectClientRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateFileProtectClientRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileProtectClientRuleResponseBody) SetId(v int64) *CreateFileProtectClientRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFileProtectClientRuleResponseBody) SetRequestId(v string) *CreateFileProtectClientRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileProtectClientRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
