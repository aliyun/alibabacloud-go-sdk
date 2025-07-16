// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelFeatureFGInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetModelFeatureFGInfoResponseBody
	GetContent() *string
	SetRequestId(v string) *GetModelFeatureFGInfoResponseBody
	GetRequestId() *string
}

type GetModelFeatureFGInfoResponseBody struct {
	// example:
	//
	// {"features": [{"feature_name": "item_id","feature_type": "id_feature","value_type": "String","expression": "item:item_id","default_value": "-1024","combiner": "mean","need_prefix": false},{"feature_name": "f1","feature_type": "lookup_feature","value_type": "Integer","map": "item:f1","key": "user:1","default_value": "0","combiner": "mean","need_prefix": false,"needDiscrete": false,"needWeighting": false,"needKey": false}],"reserves": ["f1"]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 6B662A64-E4BF-56F8-BF5F-4C63F34EC0A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModelFeatureFGInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelFeatureFGInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelFeatureFGInfoResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetModelFeatureFGInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelFeatureFGInfoResponseBody) SetContent(v string) *GetModelFeatureFGInfoResponseBody {
	s.Content = &v
	return s
}

func (s *GetModelFeatureFGInfoResponseBody) SetRequestId(v string) *GetModelFeatureFGInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelFeatureFGInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
