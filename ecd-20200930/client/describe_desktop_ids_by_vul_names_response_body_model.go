// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopIdsByVulNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopItems(v []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) *DescribeDesktopIdsByVulNamesResponseBody
	GetDesktopItems() []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems
	SetRequestId(v string) *DescribeDesktopIdsByVulNamesResponseBody
	GetRequestId() *string
}

type DescribeDesktopIdsByVulNamesResponseBody struct {
	// The array of cloud desktop information.
	DesktopItems []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems `json:"DesktopItems,omitempty" xml:"DesktopItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopIdsByVulNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) GetDesktopItems() []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	return s.DesktopItems
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) SetDesktopItems(v []*DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) *DescribeDesktopIdsByVulNamesResponseBody {
	s.DesktopItems = v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) SetRequestId(v string) *DescribeDesktopIdsByVulNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBody) Validate() error {
	if s.DesktopItems != nil {
		for _, item := range s.DesktopItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopIdsByVulNamesResponseBodyDesktopItems struct {
	// The cloud desktop ID.
	//
	// example:
	//
	// ecd-138dsptkrt00u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud desktop name.
	//
	// example:
	//
	// Graphics-21
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
}

func (s DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) GoString() string {
	return s.String()
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) SetDesktopId(v string) *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) SetDesktopName(v string) *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopIdsByVulNamesResponseBodyDesktopItems) Validate() error {
	return dara.Validate(s)
}
