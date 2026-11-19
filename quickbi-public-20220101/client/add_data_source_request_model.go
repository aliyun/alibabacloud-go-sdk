// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddModel(v string) *AddDataSourceRequest
	GetAddModel() *string
}

type AddDataSourceRequest struct {
	// You can refer to the logic of adding a data source on the Quick BI data source page. When you click **Test Connection**, the `restapi/datasource/detect` API operation is triggered. Simulate the request payload of this API operation by passing a JSON object of the same model, along with a valid `userId` and workspace ID in Quick BI. In version 7.1, the `userAsOwner` parameter is added. If you set this parameter to true, the data source is created as the specified `userId`.
	//
	// <notice>A small number of data sources are not supported. If the parameters are the same as the request payload of the test API operation but the data source still fails to be added, the current data source type does not support creation through this API operation.</notice>
	//
	// <notice>Do not pass the `encode` parameter. Currently, this API operation does not support encryption mode for creating data sources, and authentication that involves uploading files is not supported.</notice>
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "userId": "235345xxxxx24336",
	//
	//   "userAsOwner":true,
	//
	//   "workspaceId": "235345xxxxx24336",
	//
	//   "dsType": "postgresql",
	//
	//   "config": {
	//
	//     "category": "local",
	//
	//     "customDriverId": null,
	//
	//     "ssl": false,
	//
	//     "uploadFile": true,
	//
	//     "sslConfig": null
	//
	//   },
	//
	//   "showName": "test",
	//
	//   "address": "test",
	//
	//   "port": "5432",
	//
	//   "instance": "postgres",
	//
	//   "schema": "public",
	//
	//   "userName": "13****34",
	//
	//   "password": "12****425",
	//
	// }
	AddModel *string `json:"AddModel,omitempty" xml:"AddModel,omitempty"`
}

func (s AddDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceRequest) GoString() string {
	return s.String()
}

func (s *AddDataSourceRequest) GetAddModel() *string {
	return s.AddModel
}

func (s *AddDataSourceRequest) SetAddModel(v string) *AddDataSourceRequest {
	s.AddModel = &v
	return s
}

func (s *AddDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
