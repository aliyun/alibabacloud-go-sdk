// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKBSyncLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLinkId(v string) *UpdateKBSyncLinkResponseBody
	GetLinkId() *string
	SetRequestId(v string) *UpdateKBSyncLinkResponseBody
	GetRequestId() *string
	SetSyncSchedule(v string) *UpdateKBSyncLinkResponseBody
	GetSyncSchedule() *string
}

type UpdateKBSyncLinkResponseBody struct {
	// example:
	//
	// pkbl-xxxxx
	LinkId *string `json:"LinkId,omitempty" xml:"LinkId,omitempty"`
	// example:
	//
	// EBEAA83D-1734-42E3-85E3-E25F6E******
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SyncSchedule *string `json:"SyncSchedule,omitempty" xml:"SyncSchedule,omitempty"`
}

func (s UpdateKBSyncLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKBSyncLinkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKBSyncLinkResponseBody) GetLinkId() *string {
	return s.LinkId
}

func (s *UpdateKBSyncLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKBSyncLinkResponseBody) GetSyncSchedule() *string {
	return s.SyncSchedule
}

func (s *UpdateKBSyncLinkResponseBody) SetLinkId(v string) *UpdateKBSyncLinkResponseBody {
	s.LinkId = &v
	return s
}

func (s *UpdateKBSyncLinkResponseBody) SetRequestId(v string) *UpdateKBSyncLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKBSyncLinkResponseBody) SetSyncSchedule(v string) *UpdateKBSyncLinkResponseBody {
	s.SyncSchedule = &v
	return s
}

func (s *UpdateKBSyncLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
