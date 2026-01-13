// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecallManagementTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementTableId(v string) *CreateRecallManagementTableResponseBody
	GetRecallManagementTableId() *string
	SetRequestId(v string) *CreateRecallManagementTableResponseBody
	GetRequestId() *string
}

type CreateRecallManagementTableResponseBody struct {
	// example:
	//
	// 3
	RecallManagementTableId *string `json:"RecallManagementTableId,omitempty" xml:"RecallManagementTableId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecallManagementTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecallManagementTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecallManagementTableResponseBody) GetRecallManagementTableId() *string {
	return s.RecallManagementTableId
}

func (s *CreateRecallManagementTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecallManagementTableResponseBody) SetRecallManagementTableId(v string) *CreateRecallManagementTableResponseBody {
	s.RecallManagementTableId = &v
	return s
}

func (s *CreateRecallManagementTableResponseBody) SetRequestId(v string) *CreateRecallManagementTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecallManagementTableResponseBody) Validate() error {
	return dara.Validate(s)
}
