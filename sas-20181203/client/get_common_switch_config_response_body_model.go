// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommonSwitchConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetCommonSwitchConfigResponseBodyData) *GetCommonSwitchConfigResponseBody
	GetData() *GetCommonSwitchConfigResponseBodyData
	SetRequestId(v string) *GetCommonSwitchConfigResponseBody
	GetRequestId() *string
}

type GetCommonSwitchConfigResponseBody struct {
	// The data returned.
	Data *GetCommonSwitchConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCommonSwitchConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommonSwitchConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommonSwitchConfigResponseBody) GetData() *GetCommonSwitchConfigResponseBodyData {
	return s.Data
}

func (s *GetCommonSwitchConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommonSwitchConfigResponseBody) SetData(v *GetCommonSwitchConfigResponseBodyData) *GetCommonSwitchConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetCommonSwitchConfigResponseBody) SetRequestId(v string) *GetCommonSwitchConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommonSwitchConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCommonSwitchConfigResponseBodyData struct {
	// Specifies whether to turn on the switch for newly added servers. Valid values:
	//
	// 	- **add**: By default, the switch is turned on for newly added servers.
	//
	// 	- **del**: By default, the switch is turned off for newly added servers.
	//
	// example:
	//
	// add
	TargetDefault *string `json:"TargetDefault,omitempty" xml:"TargetDefault,omitempty"`
	// The status of the synchronization. Valid values:
	//
	// 	- **sync**: The modifications are being synchronized.
	//
	// 	- **valid**: The modifications has taken effect.
	//
	// example:
	//
	// valid
	TargetSyncStatus *string `json:"TargetSyncStatus,omitempty" xml:"TargetSyncStatus,omitempty"`
}

func (s GetCommonSwitchConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCommonSwitchConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommonSwitchConfigResponseBodyData) GetTargetDefault() *string {
	return s.TargetDefault
}

func (s *GetCommonSwitchConfigResponseBodyData) GetTargetSyncStatus() *string {
	return s.TargetSyncStatus
}

func (s *GetCommonSwitchConfigResponseBodyData) SetTargetDefault(v string) *GetCommonSwitchConfigResponseBodyData {
	s.TargetDefault = &v
	return s
}

func (s *GetCommonSwitchConfigResponseBodyData) SetTargetSyncStatus(v string) *GetCommonSwitchConfigResponseBodyData {
	s.TargetSyncStatus = &v
	return s
}

func (s *GetCommonSwitchConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
