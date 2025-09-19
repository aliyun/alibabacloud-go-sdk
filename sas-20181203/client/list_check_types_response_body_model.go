// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCheckTypesResponseBody
	GetCode() *string
	SetCount(v int32) *ListCheckTypesResponseBody
	GetCount() *int32
	SetData(v []*ListCheckTypesResponseBodyData) *ListCheckTypesResponseBody
	GetData() []*ListCheckTypesResponseBodyData
	SetHttpStatusCode(v int32) *ListCheckTypesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCheckTypesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCheckTypesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCheckTypesResponseBody
	GetSuccess() *bool
}

type ListCheckTypesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data returned.
	Data []*ListCheckTypesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// D65AADFC-1D20-5A6A-8F6A-9FA53C0D****
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

func (s ListCheckTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckTypesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCheckTypesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListCheckTypesResponseBody) GetData() []*ListCheckTypesResponseBodyData {
	return s.Data
}

func (s *ListCheckTypesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCheckTypesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCheckTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckTypesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCheckTypesResponseBody) SetCode(v string) *ListCheckTypesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCheckTypesResponseBody) SetCount(v int32) *ListCheckTypesResponseBody {
	s.Count = &v
	return s
}

func (s *ListCheckTypesResponseBody) SetData(v []*ListCheckTypesResponseBodyData) *ListCheckTypesResponseBody {
	s.Data = v
	return s
}

func (s *ListCheckTypesResponseBody) SetHttpStatusCode(v int32) *ListCheckTypesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCheckTypesResponseBody) SetMessage(v string) *ListCheckTypesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCheckTypesResponseBody) SetRequestId(v string) *ListCheckTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckTypesResponseBody) SetSuccess(v bool) *ListCheckTypesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCheckTypesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCheckTypesResponseBodyData struct {
	// The detail of check items.
	CheckDetails []*ListCheckTypesResponseBodyDataCheckDetails `json:"CheckDetails,omitempty" xml:"CheckDetails,omitempty" type:"Repeated"`
	// The type of the check item.
	//
	// example:
	//
	// data_integrity
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The display name of the check item type.
	//
	// example:
	//
	// Data Integrity
	CheckTypeDisName *string `json:"CheckTypeDisName,omitempty" xml:"CheckTypeDisName,omitempty"`
}

func (s ListCheckTypesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCheckTypesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCheckTypesResponseBodyData) GetCheckDetails() []*ListCheckTypesResponseBodyDataCheckDetails {
	return s.CheckDetails
}

func (s *ListCheckTypesResponseBodyData) GetCheckType() *string {
	return s.CheckType
}

func (s *ListCheckTypesResponseBodyData) GetCheckTypeDisName() *string {
	return s.CheckTypeDisName
}

func (s *ListCheckTypesResponseBodyData) SetCheckDetails(v []*ListCheckTypesResponseBodyDataCheckDetails) *ListCheckTypesResponseBodyData {
	s.CheckDetails = v
	return s
}

func (s *ListCheckTypesResponseBodyData) SetCheckType(v string) *ListCheckTypesResponseBodyData {
	s.CheckType = &v
	return s
}

func (s *ListCheckTypesResponseBodyData) SetCheckTypeDisName(v string) *ListCheckTypesResponseBodyData {
	s.CheckTypeDisName = &v
	return s
}

func (s *ListCheckTypesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListCheckTypesResponseBodyDataCheckDetails struct {
	// The list of the baseline categories of attribution.
	AffiliatedRiskTypes []*string `json:"AffiliatedRiskTypes,omitempty" xml:"AffiliatedRiskTypes,omitempty" type:"Repeated"`
	// The list of baselines attribution.
	AffiliatedRisks []*string `json:"AffiliatedRisks,omitempty" xml:"AffiliatedRisks,omitempty" type:"Repeated"`
	// The ID of the check item.
	//
	// example:
	//
	// 31
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// Configure the idle session timeout period.
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
}

func (s ListCheckTypesResponseBodyDataCheckDetails) String() string {
	return dara.Prettify(s)
}

func (s ListCheckTypesResponseBodyDataCheckDetails) GoString() string {
	return s.String()
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) GetAffiliatedRiskTypes() []*string {
	return s.AffiliatedRiskTypes
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) GetAffiliatedRisks() []*string {
	return s.AffiliatedRisks
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) GetCheckItem() *string {
	return s.CheckItem
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) SetAffiliatedRiskTypes(v []*string) *ListCheckTypesResponseBodyDataCheckDetails {
	s.AffiliatedRiskTypes = v
	return s
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) SetAffiliatedRisks(v []*string) *ListCheckTypesResponseBodyDataCheckDetails {
	s.AffiliatedRisks = v
	return s
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) SetCheckId(v int64) *ListCheckTypesResponseBodyDataCheckDetails {
	s.CheckId = &v
	return s
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) SetCheckItem(v string) *ListCheckTypesResponseBodyDataCheckDetails {
	s.CheckItem = &v
	return s
}

func (s *ListCheckTypesResponseBodyDataCheckDetails) Validate() error {
	return dara.Validate(s)
}
