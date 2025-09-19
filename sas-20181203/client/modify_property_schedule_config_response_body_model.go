// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPropertyScheduleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModifyResult(v bool) *ModifyPropertyScheduleConfigResponseBody
	GetModifyResult() *bool
	SetRequestId(v string) *ModifyPropertyScheduleConfigResponseBody
	GetRequestId() *string
}

type ModifyPropertyScheduleConfigResponseBody struct {
	// Indicates whether the collection frequency of asset fingerprints is modified. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	ModifyResult *bool `json:"ModifyResult,omitempty" xml:"ModifyResult,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C0D1DF51-D879-55A8-B6C7-7D3B913EF34D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPropertyScheduleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyPropertyScheduleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPropertyScheduleConfigResponseBody) GetModifyResult() *bool {
	return s.ModifyResult
}

func (s *ModifyPropertyScheduleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyPropertyScheduleConfigResponseBody) SetModifyResult(v bool) *ModifyPropertyScheduleConfigResponseBody {
	s.ModifyResult = &v
	return s
}

func (s *ModifyPropertyScheduleConfigResponseBody) SetRequestId(v string) *ModifyPropertyScheduleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyPropertyScheduleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
