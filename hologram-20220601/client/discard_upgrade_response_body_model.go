// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscardUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DiscardUpgradeResponseBodyData) *DiscardUpgradeResponseBody
	GetData() []*DiscardUpgradeResponseBodyData
	SetErrorCode(v string) *DiscardUpgradeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DiscardUpgradeResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DiscardUpgradeResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DiscardUpgradeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DiscardUpgradeResponseBody
	GetSuccess() *bool
}

type DiscardUpgradeResponseBody struct {
	Data []*DiscardUpgradeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A47D57BE-59D1-51A6-9CC3-080C07852C4E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DiscardUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiscardUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *DiscardUpgradeResponseBody) GetData() []*DiscardUpgradeResponseBodyData {
	return s.Data
}

func (s *DiscardUpgradeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DiscardUpgradeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DiscardUpgradeResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DiscardUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiscardUpgradeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DiscardUpgradeResponseBody) SetData(v []*DiscardUpgradeResponseBodyData) *DiscardUpgradeResponseBody {
	s.Data = v
	return s
}

func (s *DiscardUpgradeResponseBody) SetErrorCode(v string) *DiscardUpgradeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DiscardUpgradeResponseBody) SetErrorMessage(v string) *DiscardUpgradeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DiscardUpgradeResponseBody) SetHttpStatusCode(v string) *DiscardUpgradeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DiscardUpgradeResponseBody) SetRequestId(v string) *DiscardUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiscardUpgradeResponseBody) SetSuccess(v bool) *DiscardUpgradeResponseBody {
	s.Success = &v
	return s
}

func (s *DiscardUpgradeResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DiscardUpgradeResponseBodyData struct {
	// example:
	//
	// null
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// example:
	//
	// null
	ReasonKey *string `json:"ReasonKey,omitempty" xml:"ReasonKey,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DiscardUpgradeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DiscardUpgradeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DiscardUpgradeResponseBodyData) GetFailReason() *string {
	return s.FailReason
}

func (s *DiscardUpgradeResponseBodyData) GetReasonKey() *string {
	return s.ReasonKey
}

func (s *DiscardUpgradeResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *DiscardUpgradeResponseBodyData) SetFailReason(v string) *DiscardUpgradeResponseBodyData {
	s.FailReason = &v
	return s
}

func (s *DiscardUpgradeResponseBodyData) SetReasonKey(v string) *DiscardUpgradeResponseBodyData {
	s.ReasonKey = &v
	return s
}

func (s *DiscardUpgradeResponseBodyData) SetSuccess(v bool) *DiscardUpgradeResponseBodyData {
	s.Success = &v
	return s
}

func (s *DiscardUpgradeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
