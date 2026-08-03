// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineAICoachScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineAICoachScriptResponseBody
	GetRequestId() *string
	SetStatus(v string) *OfflineAICoachScriptResponseBody
	GetStatus() *string
}

type OfflineAICoachScriptResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OfflineAICoachScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineAICoachScriptResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineAICoachScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineAICoachScriptResponseBody) GetStatus() *string {
	return s.Status
}

func (s *OfflineAICoachScriptResponseBody) SetRequestId(v string) *OfflineAICoachScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineAICoachScriptResponseBody) SetStatus(v string) *OfflineAICoachScriptResponseBody {
	s.Status = &v
	return s
}

func (s *OfflineAICoachScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
