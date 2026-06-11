// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDynamicTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDynamicTagResponseBody
	GetRequestId() *string
	SetResult(v []*ListDynamicTagResponseBodyResult) *ListDynamicTagResponseBody
	GetResult() []*ListDynamicTagResponseBodyResult
	SetSuccess(v bool) *ListDynamicTagResponseBody
	GetSuccess() *bool
}

type ListDynamicTagResponseBody struct {
	// example:
	//
	// D787E1A**********DF8D885
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListDynamicTagResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDynamicTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDynamicTagResponseBody) GetResult() []*ListDynamicTagResponseBodyResult {
	return s.Result
}

func (s *ListDynamicTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDynamicTagResponseBody) SetRequestId(v string) *ListDynamicTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicTagResponseBody) SetResult(v []*ListDynamicTagResponseBodyResult) *ListDynamicTagResponseBody {
	s.Result = v
	return s
}

func (s *ListDynamicTagResponseBody) SetSuccess(v bool) *ListDynamicTagResponseBody {
	s.Success = &v
	return s
}

func (s *ListDynamicTagResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDynamicTagResponseBodyResult struct {
	// example:
	//
	// site_id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// cfg****14352318681088
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// eip
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// example:
	//
	// a201c85c-******
	DsId *string `json:"DsId,omitempty" xml:"DsId,omitempty"`
	// example:
	//
	// 2fe4fbd8-****-****-b3e1-e92c7af083ea
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// example:
	//
	// 1
	RelatedAttribute *int32 `json:"RelatedAttribute,omitempty" xml:"RelatedAttribute,omitempty"`
	// example:
	//
	// testTable02\\"\\"
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDynamicTagResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDynamicTagResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDynamicTagResponseBodyResult) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDynamicTagResponseBodyResult) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListDynamicTagResponseBodyResult) GetConfigName() *string {
	return s.ConfigName
}

func (s *ListDynamicTagResponseBodyResult) GetDsId() *string {
	return s.DsId
}

func (s *ListDynamicTagResponseBodyResult) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListDynamicTagResponseBodyResult) GetRelatedAttribute() *int32 {
	return s.RelatedAttribute
}

func (s *ListDynamicTagResponseBodyResult) GetTableName() *string {
	return s.TableName
}

func (s *ListDynamicTagResponseBodyResult) SetColumnName(v string) *ListDynamicTagResponseBodyResult {
	s.ColumnName = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) SetConfigId(v string) *ListDynamicTagResponseBodyResult {
	s.ConfigId = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) SetConfigName(v string) *ListDynamicTagResponseBodyResult {
	s.ConfigName = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) SetDsId(v string) *ListDynamicTagResponseBodyResult {
	s.DsId = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) SetOrganizationId(v string) *ListDynamicTagResponseBodyResult {
	s.OrganizationId = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) SetRelatedAttribute(v int32) *ListDynamicTagResponseBodyResult {
	s.RelatedAttribute = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) SetTableName(v string) *ListDynamicTagResponseBodyResult {
	s.TableName = &v
	return s
}

func (s *ListDynamicTagResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
