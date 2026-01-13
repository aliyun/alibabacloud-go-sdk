// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceVersionConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementServiceVersionConfigId(v string) *CreateRecallManagementServiceVersionConfigResponseBody
	GetRecallManagementServiceVersionConfigId() *string
	SetRequestId(v string) *CreateRecallManagementServiceVersionConfigResponseBody
	GetRequestId() *string
}

type CreateRecallManagementServiceVersionConfigResponseBody struct {
	// example:
	//
	// 1
	RecallManagementServiceVersionConfigId *string `json:"RecallManagementServiceVersionConfigId,omitempty" xml:"RecallManagementServiceVersionConfigId,omitempty"`
	// example:
	//
	// F8F613A9-DF1C-551A-88E1-397A3981A785
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecallManagementServiceVersionConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionConfigResponseBody) GetRecallManagementServiceVersionConfigId() *string {
	return s.RecallManagementServiceVersionConfigId
}

func (s *CreateRecallManagementServiceVersionConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecallManagementServiceVersionConfigResponseBody) SetRecallManagementServiceVersionConfigId(v string) *CreateRecallManagementServiceVersionConfigResponseBody {
	s.RecallManagementServiceVersionConfigId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigResponseBody) SetRequestId(v string) *CreateRecallManagementServiceVersionConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
