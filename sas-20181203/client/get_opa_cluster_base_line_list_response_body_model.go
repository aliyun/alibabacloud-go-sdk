// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterBaseLineListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpaClusterBaseLineListResponseBody
	GetCode() *string
	SetData(v []*GetOpaClusterBaseLineListResponseBodyData) *GetOpaClusterBaseLineListResponseBody
	GetData() []*GetOpaClusterBaseLineListResponseBodyData
	SetMessage(v string) *GetOpaClusterBaseLineListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpaClusterBaseLineListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOpaClusterBaseLineListResponseBody
	GetSuccess() *bool
}

type GetOpaClusterBaseLineListResponseBody struct {
	// The response code. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about baselines.
	Data []*GetOpaClusterBaseLineListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C2F2A020-1CAB-5F52-8CAF-B2ACDDFAC247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOpaClusterBaseLineListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterBaseLineListResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpaClusterBaseLineListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpaClusterBaseLineListResponseBody) GetData() []*GetOpaClusterBaseLineListResponseBodyData {
	return s.Data
}

func (s *GetOpaClusterBaseLineListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpaClusterBaseLineListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpaClusterBaseLineListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOpaClusterBaseLineListResponseBody) SetCode(v string) *GetOpaClusterBaseLineListResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBody) SetData(v []*GetOpaClusterBaseLineListResponseBodyData) *GetOpaClusterBaseLineListResponseBody {
	s.Data = v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBody) SetMessage(v string) *GetOpaClusterBaseLineListResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBody) SetRequestId(v string) *GetOpaClusterBaseLineListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBody) SetSuccess(v bool) *GetOpaClusterBaseLineListResponseBody {
	s.Success = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBody) Validate() error {
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

type GetOpaClusterBaseLineListResponseBodyData struct {
	// The alias of the baseline.
	//
	// example:
	//
	// Make sure there are no duplicate usernames or UIDs
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The key of the baseline type.
	//
	// example:
	//
	// identification
	ClassKey *string `json:"ClassKey,omitempty" xml:"ClassKey,omitempty"`
	// The key of the name for the baseline check item.
	//
	// example:
	//
	// duplicate_uid
	ItemKey *string `json:"ItemKey,omitempty" xml:"ItemKey,omitempty"`
	// The key of the name for the baseline.
	//
	// example:
	//
	// identification
	NameKey *string `json:"NameKey,omitempty" xml:"NameKey,omitempty"`
}

func (s GetOpaClusterBaseLineListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterBaseLineListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOpaClusterBaseLineListResponseBodyData) GetAlias() *string {
	return s.Alias
}

func (s *GetOpaClusterBaseLineListResponseBodyData) GetClassKey() *string {
	return s.ClassKey
}

func (s *GetOpaClusterBaseLineListResponseBodyData) GetItemKey() *string {
	return s.ItemKey
}

func (s *GetOpaClusterBaseLineListResponseBodyData) GetNameKey() *string {
	return s.NameKey
}

func (s *GetOpaClusterBaseLineListResponseBodyData) SetAlias(v string) *GetOpaClusterBaseLineListResponseBodyData {
	s.Alias = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBodyData) SetClassKey(v string) *GetOpaClusterBaseLineListResponseBodyData {
	s.ClassKey = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBodyData) SetItemKey(v string) *GetOpaClusterBaseLineListResponseBodyData {
	s.ItemKey = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBodyData) SetNameKey(v string) *GetOpaClusterBaseLineListResponseBodyData {
	s.NameKey = &v
	return s
}

func (s *GetOpaClusterBaseLineListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
