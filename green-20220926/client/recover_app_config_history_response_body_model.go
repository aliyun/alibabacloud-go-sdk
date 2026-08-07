// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverAppConfigHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RecoverAppConfigHistoryResponseBody
	GetData() *bool
	SetRequestId(v string) *RecoverAppConfigHistoryResponseBody
	GetRequestId() *string
}

type RecoverAppConfigHistoryResponseBody struct {
	// The returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverAppConfigHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverAppConfigHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverAppConfigHistoryResponseBody) GetData() *bool {
	return s.Data
}

func (s *RecoverAppConfigHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverAppConfigHistoryResponseBody) SetData(v bool) *RecoverAppConfigHistoryResponseBody {
	s.Data = &v
	return s
}

func (s *RecoverAppConfigHistoryResponseBody) SetRequestId(v string) *RecoverAppConfigHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverAppConfigHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}
