// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAICoachTaskSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseAICoachTaskSessionResponseBody
	GetRequestId() *string
	SetStatus(v string) *CloseAICoachTaskSessionResponseBody
	GetStatus() *string
}

type CloseAICoachTaskSessionResponseBody struct {
	// example:
	//
	// 0E06E0AA-D5B6-538C-8CE9-BAB79C68B690
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CloseAICoachTaskSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseAICoachTaskSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseAICoachTaskSessionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CloseAICoachTaskSessionResponseBody) SetRequestId(v string) *CloseAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseAICoachTaskSessionResponseBody) SetStatus(v string) *CloseAICoachTaskSessionResponseBody {
	s.Status = &v
	return s
}

func (s *CloseAICoachTaskSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
