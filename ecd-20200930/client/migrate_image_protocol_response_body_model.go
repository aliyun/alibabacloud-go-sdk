// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateImageProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedIds(v []*string) *MigrateImageProtocolResponseBody
	GetFailedIds() []*string
	SetRequestId(v string) *MigrateImageProtocolResponseBody
	GetRequestId() *string
}

type MigrateImageProtocolResponseBody struct {
	// The IDs of the images whose protocols failed to be updated.
	FailedIds []*string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4D4E5AF5-DF28-5FE7-85C7-9F98315B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MigrateImageProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MigrateImageProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *MigrateImageProtocolResponseBody) GetFailedIds() []*string {
	return s.FailedIds
}

func (s *MigrateImageProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MigrateImageProtocolResponseBody) SetFailedIds(v []*string) *MigrateImageProtocolResponseBody {
	s.FailedIds = v
	return s
}

func (s *MigrateImageProtocolResponseBody) SetRequestId(v string) *MigrateImageProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *MigrateImageProtocolResponseBody) Validate() error {
	return dara.Validate(s)
}
