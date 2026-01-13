// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishRecallManagementTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecallManagementJobId(v string) *PublishRecallManagementTableResponseBody
	GetRecallManagementJobId() *string
	SetRequestId(v string) *PublishRecallManagementTableResponseBody
	GetRequestId() *string
}

type PublishRecallManagementTableResponseBody struct {
	// example:
	//
	// 1
	RecallManagementJobId *string `json:"RecallManagementJobId,omitempty" xml:"RecallManagementJobId,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishRecallManagementTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishRecallManagementTableResponseBody) GoString() string {
	return s.String()
}

func (s *PublishRecallManagementTableResponseBody) GetRecallManagementJobId() *string {
	return s.RecallManagementJobId
}

func (s *PublishRecallManagementTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishRecallManagementTableResponseBody) SetRecallManagementJobId(v string) *PublishRecallManagementTableResponseBody {
	s.RecallManagementJobId = &v
	return s
}

func (s *PublishRecallManagementTableResponseBody) SetRequestId(v string) *PublishRecallManagementTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishRecallManagementTableResponseBody) Validate() error {
	return dara.Validate(s)
}
