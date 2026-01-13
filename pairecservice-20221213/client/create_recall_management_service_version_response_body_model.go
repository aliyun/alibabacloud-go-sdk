// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementServiceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementServiceVersionId(v string) *CreateRecallManagementServiceVersionResponseBody
	GetRecallManagementServiceVersionId() *string
	SetRequestId(v string) *CreateRecallManagementServiceVersionResponseBody
	GetRequestId() *string
}

type CreateRecallManagementServiceVersionResponseBody struct {
	// example:
	//
	// 1
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecallManagementServiceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementServiceVersionResponseBody) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *CreateRecallManagementServiceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecallManagementServiceVersionResponseBody) SetRecallManagementServiceVersionId(v string) *CreateRecallManagementServiceVersionResponseBody {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionResponseBody) SetRequestId(v string) *CreateRecallManagementServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecallManagementServiceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
