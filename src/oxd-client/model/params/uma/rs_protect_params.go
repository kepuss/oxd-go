//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package uma

import "oxd-client/model/params/uma/protect"

type RsProtectRequestParams struct {
	OxdId string `json:"oxd_id"`
	Resources []protect.RsResource `json:"resources"`
}

type RsProtectResponseParams struct {
	OxdId string `json:"oxd_id"`
}
