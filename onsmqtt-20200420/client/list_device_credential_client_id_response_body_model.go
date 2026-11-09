// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeviceCredentialClientIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceCredentialClientIdList(v *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) *ListDeviceCredentialClientIdResponseBody
	GetDeviceCredentialClientIdList() *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList
	SetRequestId(v string) *ListDeviceCredentialClientIdResponseBody
	GetRequestId() *string
}

type ListDeviceCredentialClientIdResponseBody struct {
	// Returns the information list.
	DeviceCredentialClientIdList *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList `json:"DeviceCredentialClientIdList,omitempty" xml:"DeviceCredentialClientIdList,omitempty" type:"Struct"`
	// Public parameters, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeviceCredentialClientIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCredentialClientIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdResponseBody) GetDeviceCredentialClientIdList() *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	return s.DeviceCredentialClientIdList
}

func (s *ListDeviceCredentialClientIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeviceCredentialClientIdResponseBody) SetDeviceCredentialClientIdList(v *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) *ListDeviceCredentialClientIdResponseBody {
	s.DeviceCredentialClientIdList = v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBody) SetRequestId(v string) *ListDeviceCredentialClientIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBody) Validate() error {
	if s.DeviceCredentialClientIdList != nil {
		if err := s.DeviceCredentialClientIdList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList struct {
	// Client list.
	ClientIdList []*string `json:"ClientIdList,omitempty" xml:"ClientIdList,omitempty" type:"Repeated"`
	// Indicates whether there is a token (Token) for the next query. Values:
	//
	// - For the first query and when there is no next query, this field does not need to be filled.
	//
	// - If there is a next query, the value should be the NextToken returned from the previous API call.
	//
	// example:
	//
	// 634dxxxxx75b5f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number of the returned query records.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of results to display per page.
	//
	// example:
	//
	// 100
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of query results.
	//
	// example:
	//
	// 10
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) String() string {
	return dara.Prettify(s)
}

func (s ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GoString() string {
	return s.String()
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GetClientIdList() []*string {
	return s.ClientIdList
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GetPageNo() *string {
	return s.PageNo
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) GetTotal() *string {
	return s.Total
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetClientIdList(v []*string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.ClientIdList = v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetNextToken(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.NextToken = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetPageNo(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.PageNo = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetPageSize(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.PageSize = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) SetTotal(v string) *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList {
	s.Total = &v
	return s
}

func (s *ListDeviceCredentialClientIdResponseBodyDeviceCredentialClientIdList) Validate() error {
	return dara.Validate(s)
}
