// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRCKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SyncRCKeyPairResponseBodyData) *SyncRCKeyPairResponseBody
	GetData() *SyncRCKeyPairResponseBodyData
	SetRequestId(v string) *SyncRCKeyPairResponseBody
	GetRequestId() *string
}

type SyncRCKeyPairResponseBody struct {
	Data      *SyncRCKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SyncRCKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncRCKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *SyncRCKeyPairResponseBody) GetData() *SyncRCKeyPairResponseBodyData {
	return s.Data
}

func (s *SyncRCKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncRCKeyPairResponseBody) SetData(v *SyncRCKeyPairResponseBodyData) *SyncRCKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *SyncRCKeyPairResponseBody) SetRequestId(v string) *SyncRCKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncRCKeyPairResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncRCKeyPairResponseBodyData struct {
	IsSyncInfo *bool `json:"IsSyncInfo,omitempty" xml:"IsSyncInfo,omitempty"`
}

func (s SyncRCKeyPairResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SyncRCKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *SyncRCKeyPairResponseBodyData) GetIsSyncInfo() *bool {
	return s.IsSyncInfo
}

func (s *SyncRCKeyPairResponseBodyData) SetIsSyncInfo(v bool) *SyncRCKeyPairResponseBodyData {
	s.IsSyncInfo = &v
	return s
}

func (s *SyncRCKeyPairResponseBodyData) Validate() error {
	return dara.Validate(s)
}
