// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupItems(v *DescribeClusterWhiteListResponseBodyGroupItems) *DescribeClusterWhiteListResponseBody
	GetGroupItems() *DescribeClusterWhiteListResponseBodyGroupItems
	SetRequestId(v string) *DescribeClusterWhiteListResponseBody
	GetRequestId() *string
}

type DescribeClusterWhiteListResponseBody struct {
	GroupItems *DescribeClusterWhiteListResponseBodyGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Struct"`
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListResponseBody) GetGroupItems() *DescribeClusterWhiteListResponseBodyGroupItems {
	return s.GroupItems
}

func (s *DescribeClusterWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterWhiteListResponseBody) SetGroupItems(v *DescribeClusterWhiteListResponseBodyGroupItems) *DescribeClusterWhiteListResponseBody {
	s.GroupItems = v
	return s
}

func (s *DescribeClusterWhiteListResponseBody) SetRequestId(v string) *DescribeClusterWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterWhiteListResponseBody) Validate() error {
	if s.GroupItems != nil {
		if err := s.GroupItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterWhiteListResponseBodyGroupItems struct {
	WhiteIp []*string `json:"WhiteIp,omitempty" xml:"WhiteIp,omitempty" type:"Repeated"`
}

func (s DescribeClusterWhiteListResponseBodyGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterWhiteListResponseBodyGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeClusterWhiteListResponseBodyGroupItems) GetWhiteIp() []*string {
	return s.WhiteIp
}

func (s *DescribeClusterWhiteListResponseBodyGroupItems) SetWhiteIp(v []*string) *DescribeClusterWhiteListResponseBodyGroupItems {
	s.WhiteIp = v
	return s
}

func (s *DescribeClusterWhiteListResponseBodyGroupItems) Validate() error {
	return dara.Validate(s)
}
