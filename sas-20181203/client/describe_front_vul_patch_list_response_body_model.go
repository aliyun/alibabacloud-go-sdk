// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrontVulPatchListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFrontPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchList) *DescribeFrontVulPatchListResponseBody
	GetFrontPatchList() []*DescribeFrontVulPatchListResponseBodyFrontPatchList
	SetRequestId(v string) *DescribeFrontVulPatchListResponseBody
	GetRequestId() *string
}

type DescribeFrontVulPatchListResponseBody struct {
	// An array consisting of the information about the pre-patches that are required to fix the specified Windows system vulnerability.
	FrontPatchList []*DescribeFrontVulPatchListResponseBodyFrontPatchList `json:"FrontPatchList,omitempty" xml:"FrontPatchList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F929E952-EBFC-56C3-BD35-BF8B59024C68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFrontVulPatchListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBody) GetFrontPatchList() []*DescribeFrontVulPatchListResponseBodyFrontPatchList {
	return s.FrontPatchList
}

func (s *DescribeFrontVulPatchListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFrontVulPatchListResponseBody) SetFrontPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchList) *DescribeFrontVulPatchListResponseBody {
	s.FrontPatchList = v
	return s
}

func (s *DescribeFrontVulPatchListResponseBody) SetRequestId(v string) *DescribeFrontVulPatchListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBody) Validate() error {
	if s.FrontPatchList != nil {
		for _, item := range s.FrontPatchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFrontVulPatchListResponseBodyFrontPatchList struct {
	// An array consisting of the pre-patches that are required to fix the specified Windows system vulnerability on the server.
	PatchList []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList `json:"PatchList,omitempty" xml:"PatchList,omitempty" type:"Repeated"`
	// The UUID of the server.
	//
	// example:
	//
	// 1587bedb-fdb4-48c4-9330-4545****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) GetPatchList() []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	return s.PatchList
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) SetPatchList(v []*DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) *DescribeFrontVulPatchListResponseBodyFrontPatchList {
	s.PatchList = v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) SetUuid(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchList {
	s.Uuid = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchList) Validate() error {
	if s.PatchList != nil {
		for _, item := range s.PatchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList struct {
	// The name of the vulnerability.
	//
	// example:
	//
	// RHBA-2019:2599: krb5 bug fix update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The version number of the pre-patch that is required to fix the Windows system vulnerability.
	//
	// example:
	//
	// 4523204
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) GetName() *string {
	return s.Name
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) SetAliasName(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	s.AliasName = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) SetName(v string) *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList {
	s.Name = &v
	return s
}

func (s *DescribeFrontVulPatchListResponseBodyFrontPatchListPatchList) Validate() error {
	return dara.Validate(s)
}
