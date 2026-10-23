// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterDataInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v string) *GetClusterDataInformationRequest
	GetBody() *string
}

type GetClusterDataInformationRequest struct {
	// The cluster type, access endpoint, and index to be queried.
	//
	// example:
	//
	// {"dataSourceType":"elasticsearch","endpoint":"http://es-cn-npk1shyiq000d****.public.elasticsearch.aliyuncs.com:9200","username":"elastic","password":"xxxxxx","index":"default","type":"default"}
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterDataInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClusterDataInformationRequest) GoString() string {
	return s.String()
}

func (s *GetClusterDataInformationRequest) GetBody() *string {
	return s.Body
}

func (s *GetClusterDataInformationRequest) SetBody(v string) *GetClusterDataInformationRequest {
	s.Body = &v
	return s
}

func (s *GetClusterDataInformationRequest) Validate() error {
	return dara.Validate(s)
}
