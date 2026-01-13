// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementServiceId(v string) *CreateRecallManagementServiceResponseBody
	GetRecallManagementServiceId() *string
	SetRecallManagementServiceVersionId(v string) *CreateRecallManagementServiceResponseBody
	GetRecallManagementServiceVersionId() *string
	SetRequestId(v string) *CreateRecallManagementServiceResponseBody
	GetRequestId() *string
}

type CreateRecallManagementServiceResponseBody struct {
	// example:
	//
	// 1
	RecallManagementServiceId *string `json:"RecallManagementServiceId,omitempty" xml:"RecallManagementServiceId,omitempty"`
	// example:
	//
	// 3
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceResponseBody) GetRecallManagementServiceId() *string {
	return s.RecallManagementServiceId
}

func (s *CreateRecallManagementServiceResponseBody) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *CreateRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecallManagementServiceResponseBody) SetRecallManagementServiceId(v string) *CreateRecallManagementServiceResponseBody {
	s.RecallManagementServiceId = &v
	return s
}

func (s *CreateRecallManagementServiceResponseBody) SetRecallManagementServiceVersionId(v string) *CreateRecallManagementServiceResponseBody {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *CreateRecallManagementServiceResponseBody) SetRequestId(v string) *CreateRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
