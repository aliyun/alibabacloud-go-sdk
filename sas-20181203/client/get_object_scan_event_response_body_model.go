// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetObjectScanEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetObjectScanEventResponseBodyData) *GetObjectScanEventResponseBody
	GetData() *GetObjectScanEventResponseBodyData
	SetRequestId(v string) *GetObjectScanEventResponseBody
	GetRequestId() *string
}

type GetObjectScanEventResponseBody struct {
	// The response parameters.
	Data *GetObjectScanEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 30CBF632-109F-596F-97F2-451C8B2A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetObjectScanEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetObjectScanEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectScanEventResponseBody) GetData() *GetObjectScanEventResponseBodyData {
	return s.Data
}

func (s *GetObjectScanEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetObjectScanEventResponseBody) SetData(v *GetObjectScanEventResponseBodyData) *GetObjectScanEventResponseBody {
	s.Data = v
	return s
}

func (s *GetObjectScanEventResponseBody) SetRequestId(v string) *GetObjectScanEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetObjectScanEventResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetObjectScanEventResponseBodyData struct {
	// The details of the alert event.
	Details []*GetObjectScanEventResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The name of the alert item.
	//
	// example:
	//
	// WebShell
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The name of the object.
	//
	// example:
	//
	// sca_2023****
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The MD5 hash value of the object.
	//
	// example:
	//
	// 0552c44e243abdea1729d4507bce****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
}

func (s GetObjectScanEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetObjectScanEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetObjectScanEventResponseBodyData) GetDetails() []*GetObjectScanEventResponseBodyDataDetails {
	return s.Details
}

func (s *GetObjectScanEventResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *GetObjectScanEventResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetObjectScanEventResponseBodyData) GetMd5() *string {
	return s.Md5
}

func (s *GetObjectScanEventResponseBodyData) SetDetails(v []*GetObjectScanEventResponseBodyDataDetails) *GetObjectScanEventResponseBodyData {
	s.Details = v
	return s
}

func (s *GetObjectScanEventResponseBodyData) SetEventName(v string) *GetObjectScanEventResponseBodyData {
	s.EventName = &v
	return s
}

func (s *GetObjectScanEventResponseBodyData) SetFileName(v string) *GetObjectScanEventResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetObjectScanEventResponseBodyData) SetMd5(v string) *GetObjectScanEventResponseBodyData {
	s.Md5 = &v
	return s
}

func (s *GetObjectScanEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetObjectScanEventResponseBodyDataDetails struct {
	// The type of the item.
	//
	// example:
	//
	// html
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// The name of the item.
	//
	// example:
	//
	// DownloadUrl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The display name of the item.
	//
	// example:
	//
	// DownloadUrl
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// The type of the item.
	//
	// example:
	//
	// html
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the item.
	//
	// example:
	//
	// http://****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The display value of the item.
	//
	// example:
	//
	// http://****
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s GetObjectScanEventResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s GetObjectScanEventResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *GetObjectScanEventResponseBodyDataDetails) GetInfoType() *string {
	return s.InfoType
}

func (s *GetObjectScanEventResponseBodyDataDetails) GetName() *string {
	return s.Name
}

func (s *GetObjectScanEventResponseBodyDataDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *GetObjectScanEventResponseBodyDataDetails) GetType() *string {
	return s.Type
}

func (s *GetObjectScanEventResponseBodyDataDetails) GetValue() *string {
	return s.Value
}

func (s *GetObjectScanEventResponseBodyDataDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *GetObjectScanEventResponseBodyDataDetails) SetInfoType(v string) *GetObjectScanEventResponseBodyDataDetails {
	s.InfoType = &v
	return s
}

func (s *GetObjectScanEventResponseBodyDataDetails) SetName(v string) *GetObjectScanEventResponseBodyDataDetails {
	s.Name = &v
	return s
}

func (s *GetObjectScanEventResponseBodyDataDetails) SetNameDisplay(v string) *GetObjectScanEventResponseBodyDataDetails {
	s.NameDisplay = &v
	return s
}

func (s *GetObjectScanEventResponseBodyDataDetails) SetType(v string) *GetObjectScanEventResponseBodyDataDetails {
	s.Type = &v
	return s
}

func (s *GetObjectScanEventResponseBodyDataDetails) SetValue(v string) *GetObjectScanEventResponseBodyDataDetails {
	s.Value = &v
	return s
}

func (s *GetObjectScanEventResponseBodyDataDetails) SetValueDisplay(v string) *GetObjectScanEventResponseBodyDataDetails {
	s.ValueDisplay = &v
	return s
}

func (s *GetObjectScanEventResponseBodyDataDetails) Validate() error {
	return dara.Validate(s)
}
