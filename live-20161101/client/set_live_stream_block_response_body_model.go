// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveStreamBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SetLiveStreamBlockResponseBody
	GetDescription() *string
	SetRequestId(v string) *SetLiveStreamBlockResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetLiveStreamBlockResponseBody
	GetStatus() *string
}

type SetLiveStreamBlockResponseBody struct {
	// The result description.
	//
	// - Success: ok.
	//
	// - Failure: failure details.
	//
	// example:
	//
	// ok
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3be7ade8-d907-483c-b24a-0dad4595beaf
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task status. Valid values:
	//
	// - ok: succeeded.
	//
	// - fail: failed.
	//
	// > The task is considered failed if any single subtask fails to be created.
	//
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetLiveStreamBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveStreamBlockResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveStreamBlockResponseBody) GetDescription() *string {
	return s.Description
}

func (s *SetLiveStreamBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveStreamBlockResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetLiveStreamBlockResponseBody) SetDescription(v string) *SetLiveStreamBlockResponseBody {
	s.Description = &v
	return s
}

func (s *SetLiveStreamBlockResponseBody) SetRequestId(v string) *SetLiveStreamBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveStreamBlockResponseBody) SetStatus(v string) *SetLiveStreamBlockResponseBody {
	s.Status = &v
	return s
}

func (s *SetLiveStreamBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
