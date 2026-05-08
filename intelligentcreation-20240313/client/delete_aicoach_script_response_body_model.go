// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAICoachScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAICoachScriptResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteAICoachScriptResponseBody
	GetStatus() *string
}

type DeleteAICoachScriptResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteAICoachScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAICoachScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAICoachScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAICoachScriptResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteAICoachScriptResponseBody) SetRequestId(v string) *DeleteAICoachScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAICoachScriptResponseBody) SetStatus(v string) *DeleteAICoachScriptResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteAICoachScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
