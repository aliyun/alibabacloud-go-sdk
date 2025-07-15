// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdsFileShareLinksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *DescribeCdsFileShareLinksRequest
	GetCdsId() *string
	SetCreators(v []*string) *DescribeCdsFileShareLinksRequest
	GetCreators() []*string
	SetMaxResults(v int32) *DescribeCdsFileShareLinksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCdsFileShareLinksRequest
	GetNextToken() *string
	SetShareId(v string) *DescribeCdsFileShareLinksRequest
	GetShareId() *string
	SetShareName(v string) *DescribeCdsFileShareLinksRequest
	GetShareName() *string
	SetStatus(v string) *DescribeCdsFileShareLinksRequest
	GetStatus() *string
}

type DescribeCdsFileShareLinksRequest struct {
	// The ID of the cloud disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+cds-532033****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The users that create the file sharing links.
	Creators []*string `json:"Creators,omitempty" xml:"Creators,omitempty" type:"Repeated"`
	// The maximum number of resources to return. Valid values: 1 to 100. Default value: 100. The number of returned resources must be less than or equal to the specified number.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies the marker after which the returned list begins. If this parameter is not specified, all results are returned. Default value: null.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the file sharing link.
	//
	// example:
	//
	// 7JQX1Fs****
	ShareId *string `json:"ShareId,omitempty" xml:"ShareId,omitempty"`
	// The sharing name for fuzzy search.
	//
	// example:
	//
	// user
	ShareName *string `json:"ShareName,omitempty" xml:"ShareName,omitempty"`
	// The file sharing status. Valid values: ● disabled: canceled ● enabled: valid
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdsFileShareLinksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdsFileShareLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdsFileShareLinksRequest) GetCdsId() *string {
	return s.CdsId
}

func (s *DescribeCdsFileShareLinksRequest) GetCreators() []*string {
	return s.Creators
}

func (s *DescribeCdsFileShareLinksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCdsFileShareLinksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCdsFileShareLinksRequest) GetShareId() *string {
	return s.ShareId
}

func (s *DescribeCdsFileShareLinksRequest) GetShareName() *string {
	return s.ShareName
}

func (s *DescribeCdsFileShareLinksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdsFileShareLinksRequest) SetCdsId(v string) *DescribeCdsFileShareLinksRequest {
	s.CdsId = &v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) SetCreators(v []*string) *DescribeCdsFileShareLinksRequest {
	s.Creators = v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) SetMaxResults(v int32) *DescribeCdsFileShareLinksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) SetNextToken(v string) *DescribeCdsFileShareLinksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) SetShareId(v string) *DescribeCdsFileShareLinksRequest {
	s.ShareId = &v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) SetShareName(v string) *DescribeCdsFileShareLinksRequest {
	s.ShareName = &v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) SetStatus(v string) *DescribeCdsFileShareLinksRequest {
	s.Status = &v
	return s
}

func (s *DescribeCdsFileShareLinksRequest) Validate() error {
	return dara.Validate(s)
}
