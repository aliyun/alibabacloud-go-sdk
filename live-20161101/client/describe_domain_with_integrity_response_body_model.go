// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainWithIntegrityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeDomainWithIntegrityResponseBodyContent) *DescribeDomainWithIntegrityResponseBody
	GetContent() []*DescribeDomainWithIntegrityResponseBodyContent
	SetRequestId(v string) *DescribeDomainWithIntegrityResponseBody
	GetRequestId() *string
}

type DescribeDomainWithIntegrityResponseBody struct {
	// The verification information.
	Content []*DescribeDomainWithIntegrityResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainWithIntegrityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainWithIntegrityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainWithIntegrityResponseBody) GetContent() []*DescribeDomainWithIntegrityResponseBodyContent {
	return s.Content
}

func (s *DescribeDomainWithIntegrityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainWithIntegrityResponseBody) SetContent(v []*DescribeDomainWithIntegrityResponseBodyContent) *DescribeDomainWithIntegrityResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDomainWithIntegrityResponseBody) SetRequestId(v string) *DescribeDomainWithIntegrityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainWithIntegrityResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainWithIntegrityResponseBodyContent struct {
	// The column names.
	Columns []*string `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The table name.
	//
	// example:
	//
	// 1637825700000
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The subpoints.
	Points []*string `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s DescribeDomainWithIntegrityResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainWithIntegrityResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) GetColumns() []*string {
	return s.Columns
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) GetName() *string {
	return s.Name
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) GetPoints() []*string {
	return s.Points
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) SetColumns(v []*string) *DescribeDomainWithIntegrityResponseBodyContent {
	s.Columns = v
	return s
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) SetName(v string) *DescribeDomainWithIntegrityResponseBodyContent {
	s.Name = &v
	return s
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) SetPoints(v []*string) *DescribeDomainWithIntegrityResponseBodyContent {
	s.Points = v
	return s
}

func (s *DescribeDomainWithIntegrityResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
