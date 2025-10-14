// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataDownloadURLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeDataDownloadURLResponseBody
	GetCode() *int64
	SetData(v *DescribeDataDownloadURLResponseBodyData) *DescribeDataDownloadURLResponseBody
	GetData() *DescribeDataDownloadURLResponseBodyData
	SetMessage(v string) *DescribeDataDownloadURLResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDataDownloadURLResponseBody
	GetRequestId() *string
}

type DescribeDataDownloadURLResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The download URLs of data files.
	Data *DescribeDataDownloadURLResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message. Success is returned for a successful request.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataDownloadURLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDownloadURLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeDataDownloadURLResponseBody) GetData() *DescribeDataDownloadURLResponseBodyData {
	return s.Data
}

func (s *DescribeDataDownloadURLResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDataDownloadURLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataDownloadURLResponseBody) SetCode(v int64) *DescribeDataDownloadURLResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) SetData(v *DescribeDataDownloadURLResponseBodyData) *DescribeDataDownloadURLResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) SetMessage(v string) *DescribeDataDownloadURLResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) SetRequestId(v string) *DescribeDataDownloadURLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataDownloadURLResponseBodyData struct {
	// The time when the data file expires. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-12-10T03:36:27Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The file servers.
	ServerList []*DescribeDataDownloadURLResponseBodyDataServerList `json:"ServerList,omitempty" xml:"ServerList,omitempty" type:"Repeated"`
	// The download URL of the data file.
	//
	// example:
	//
	// /file/1450088842124331/97a32f2a-aa2c-436a-b19c-05b20d258618/f0313053530fc727f81b7d03fad93cd2/e4c2e8edac362acab7123654b9e73432?ak=edgepaas-innerapi-daily&ts=1611229454&sign=Yycbax%2F4OsTgm6BLoxR6lPs5gKE%3D
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeDataDownloadURLResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDownloadURLResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDataDownloadURLResponseBodyData) GetServerList() []*DescribeDataDownloadURLResponseBodyDataServerList {
	return s.ServerList
}

func (s *DescribeDataDownloadURLResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *DescribeDataDownloadURLResponseBodyData) SetExpireTime(v string) *DescribeDataDownloadURLResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyData) SetServerList(v []*DescribeDataDownloadURLResponseBodyDataServerList) *DescribeDataDownloadURLResponseBodyData {
	s.ServerList = v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyData) SetUrl(v string) *DescribeDataDownloadURLResponseBodyData {
	s.Url = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyData) Validate() error {
	if s.ServerList != nil {
		for _, item := range s.ServerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataDownloadURLResponseBodyDataServerList struct {
	// The host address of the file server.
	//
	// example:
	//
	// 1.1.1.1:8080
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-chenzhou-5
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDataDownloadURLResponseBodyDataServerList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataDownloadURLResponseBodyDataServerList) GoString() string {
	return s.String()
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) GetHost() *string {
	return s.Host
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) SetHost(v string) *DescribeDataDownloadURLResponseBodyDataServerList {
	s.Host = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) SetRegionId(v string) *DescribeDataDownloadURLResponseBodyDataServerList {
	s.RegionId = &v
	return s
}

func (s *DescribeDataDownloadURLResponseBodyDataServerList) Validate() error {
	return dara.Validate(s)
}
