// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishAICoachTaskSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FinishAICoachTaskSessionResponseBody
	GetRequestId() *string
	SetStatus(v string) *FinishAICoachTaskSessionResponseBody
	GetStatus() *string
}

type FinishAICoachTaskSessionResponseBody struct {
	// example:
	//
	// 14878724-A835-578D-9DD5-4779ADCE9221
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s FinishAICoachTaskSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FinishAICoachTaskSessionResponseBody) GoString() string {
	return s.String()
}

func (s *FinishAICoachTaskSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FinishAICoachTaskSessionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *FinishAICoachTaskSessionResponseBody) SetRequestId(v string) *FinishAICoachTaskSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *FinishAICoachTaskSessionResponseBody) SetStatus(v string) *FinishAICoachTaskSessionResponseBody {
	s.Status = &v
	return s
}

func (s *FinishAICoachTaskSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
