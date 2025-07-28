// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainPunishStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDomainPunishStatusResponseBody
	GetRequestId() *string
}

type ModifyDomainPunishStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 70E65A35-22B8-567C-B0A0-A2E9****20AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDomainPunishStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainPunishStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDomainPunishStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDomainPunishStatusResponseBody) SetRequestId(v string) *ModifyDomainPunishStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDomainPunishStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
