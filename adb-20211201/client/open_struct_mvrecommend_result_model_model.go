// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenStructMVRecommendResultModel interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratedQueriesCount(v int64) *OpenStructMVRecommendResultModel
	GetAcceleratedQueriesCount() *int64
	SetBaseTables(v []*OpenStructMvBaseTableDetailModel) *OpenStructMVRecommendResultModel
	GetBaseTables() []*OpenStructMvBaseTableDetailModel
	SetSavedScanbytes(v int64) *OpenStructMVRecommendResultModel
	GetSavedScanbytes() *int64
	SetSubquery(v string) *OpenStructMVRecommendResultModel
	GetSubquery() *string
	SetSubqueryId(v int64) *OpenStructMVRecommendResultModel
	GetSubqueryId() *int64
	SetSupportIncrementalRefresh(v bool) *OpenStructMVRecommendResultModel
	GetSupportIncrementalRefresh() *bool
}

type OpenStructMVRecommendResultModel struct {
	AcceleratedQueriesCount   *int64                              `json:"AcceleratedQueriesCount,omitempty" xml:"AcceleratedQueriesCount,omitempty"`
	BaseTables                []*OpenStructMvBaseTableDetailModel `json:"BaseTables,omitempty" xml:"BaseTables,omitempty" type:"Repeated"`
	SavedScanbytes            *int64                              `json:"SavedScanbytes,omitempty" xml:"SavedScanbytes,omitempty"`
	Subquery                  *string                             `json:"Subquery,omitempty" xml:"Subquery,omitempty"`
	SubqueryId                *int64                              `json:"SubqueryId,omitempty" xml:"SubqueryId,omitempty"`
	SupportIncrementalRefresh *bool                               `json:"SupportIncrementalRefresh,omitempty" xml:"SupportIncrementalRefresh,omitempty"`
}

func (s OpenStructMVRecommendResultModel) String() string {
	return dara.Prettify(s)
}

func (s OpenStructMVRecommendResultModel) GoString() string {
	return s.String()
}

func (s *OpenStructMVRecommendResultModel) GetAcceleratedQueriesCount() *int64 {
	return s.AcceleratedQueriesCount
}

func (s *OpenStructMVRecommendResultModel) GetBaseTables() []*OpenStructMvBaseTableDetailModel {
	return s.BaseTables
}

func (s *OpenStructMVRecommendResultModel) GetSavedScanbytes() *int64 {
	return s.SavedScanbytes
}

func (s *OpenStructMVRecommendResultModel) GetSubquery() *string {
	return s.Subquery
}

func (s *OpenStructMVRecommendResultModel) GetSubqueryId() *int64 {
	return s.SubqueryId
}

func (s *OpenStructMVRecommendResultModel) GetSupportIncrementalRefresh() *bool {
	return s.SupportIncrementalRefresh
}

func (s *OpenStructMVRecommendResultModel) SetAcceleratedQueriesCount(v int64) *OpenStructMVRecommendResultModel {
	s.AcceleratedQueriesCount = &v
	return s
}

func (s *OpenStructMVRecommendResultModel) SetBaseTables(v []*OpenStructMvBaseTableDetailModel) *OpenStructMVRecommendResultModel {
	s.BaseTables = v
	return s
}

func (s *OpenStructMVRecommendResultModel) SetSavedScanbytes(v int64) *OpenStructMVRecommendResultModel {
	s.SavedScanbytes = &v
	return s
}

func (s *OpenStructMVRecommendResultModel) SetSubquery(v string) *OpenStructMVRecommendResultModel {
	s.Subquery = &v
	return s
}

func (s *OpenStructMVRecommendResultModel) SetSubqueryId(v int64) *OpenStructMVRecommendResultModel {
	s.SubqueryId = &v
	return s
}

func (s *OpenStructMVRecommendResultModel) SetSupportIncrementalRefresh(v bool) *OpenStructMVRecommendResultModel {
	s.SupportIncrementalRefresh = &v
	return s
}

func (s *OpenStructMVRecommendResultModel) Validate() error {
	if s.BaseTables != nil {
		for _, item := range s.BaseTables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
